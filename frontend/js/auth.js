const API_URL = 'http://localhost:8080';

// Wrapper for fetch that handles Authorization and token refresh automatically
async function apiFetch(endpoint, options = {}) {
    let token = localStorage.getItem('access_token');
    
    if (!options.headers) {
        options.headers = {};
    }
    
    if (token) {
        options.headers['Authorization'] = `Bearer ${token}`;
    }
    
    if (!options.headers['Content-Type'] && !(options.body instanceof FormData)) {
        options.headers['Content-Type'] = 'application/json';
    }

    let response = await fetch(`${API_URL}${endpoint}`, options);

    // If 401 Unauthorized, try to refresh the token
    if (response.status === 401) {
        const refreshToken = localStorage.getItem('refresh_token');
        if (refreshToken) {
            // Attempt to refresh
            const refreshRes = await fetch(`${API_URL}/auth/refresh`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ refresh_token: refreshToken })
            });

            if (refreshRes.ok) {
                const data = await refreshRes.json();
                localStorage.setItem('access_token', data.access_token);
                // Retry the original request with the new token
                options.headers['Authorization'] = `Bearer ${data.access_token}`;
                response = await fetch(`${API_URL}${endpoint}`, options);
            } else {
                // Refresh failed, logout
                logout();
                return response;
            }
        } else {
            // No refresh token, logout
            logout();
        }
    }

    return response;
}

// Redirects user to login if they have no access token
function requireAuth() {
    if (!localStorage.getItem('access_token')) {
        window.location.href = 'login.html';
    }
}

async function login(email, password) {
    const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });

    const data = await response.json();
    if (response.ok) {
        localStorage.setItem('access_token', data.access_token);
        localStorage.setItem('refresh_token', data.refresh_token);
        window.location.href = 'dashboard.html';
    } else {
        throw new Error(data.error || 'Login failed');
    }
}

async function register(name, email, password) {
    const response = await fetch(`${API_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password })
    });

    const data = await response.json();
    if (response.ok) {
        return data.message; // Registration successful
    } else {
        throw new Error(data.error || 'Registration failed');
    }
}

async function logout() {
    const token = localStorage.getItem('access_token');
    if (token) {
        await fetch(`${API_URL}/auth/logout`, {
            method: 'POST',
            headers: { 'Authorization': `Bearer ${token}` }
        });
    }
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    window.location.href = 'login.html';
}

// Attach event listeners if the forms exist on the current page
document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            const errorMsg = document.getElementById('errorMsg');
            errorMsg.style.display = 'none';

            try {
                await login(email, password);
            } catch (err) {
                errorMsg.textContent = err.message;
                errorMsg.style.display = 'block';
            }
        });
    }

    const registerForm = document.getElementById('registerForm');
    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const name = document.getElementById('name').value;
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            const errorMsg = document.getElementById('errorMsg');
            const successMsg = document.getElementById('successMsg');
            
            errorMsg.style.display = 'none';
            successMsg.style.display = 'none';

            try {
                const msg = await register(name, email, password);
                successMsg.textContent = msg + " You can now login.";
                successMsg.style.display = 'block';
                registerForm.reset();
            } catch (err) {
                errorMsg.textContent = err.message;
                errorMsg.style.display = 'block';
            }
        });
    }
});
