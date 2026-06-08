// Requires auth.js to be loaded first

async function getProfile() {
    const response = await apiFetch('/users/profile');
    if (!response.ok) {
        throw new Error('Failed to fetch profile');
    }
    const data = await response.json();
    return data.user;
}

async function updateProfile(name) {
    const response = await apiFetch('/users/profile', {
        method: 'PUT',
        body: JSON.stringify({ name })
    });
    
    const data = await response.json();
    if (!response.ok) {
        throw new Error(data.error || 'Failed to update profile');
    }
    return data.message;
}

// Logic for profile.html page
document.addEventListener('DOMContentLoaded', async () => {
    // Only run if we are on the profile page
    const profileForm = document.getElementById('profileForm');
    if (!profileForm) return;

    requireAuth();

    const nameInput = document.getElementById('name');
    const emailInput = document.getElementById('email');
    const roleInput = document.getElementById('role');
    const errorMsg = document.getElementById('errorMsg');
    const successMsg = document.getElementById('successMsg');

    try {
        const user = await getProfile();
        nameInput.value = user.name;
        emailInput.value = user.email;
        roleInput.value = user.role;
    } catch (err) {
        errorMsg.textContent = err.message;
        errorMsg.style.display = 'block';
    }

    profileForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        errorMsg.style.display = 'none';
        successMsg.style.display = 'none';

        try {
            const newName = nameInput.value;
            const msg = await updateProfile(newName);
            successMsg.textContent = msg;
            successMsg.style.display = 'block';
        } catch (err) {
            errorMsg.textContent = err.message;
            errorMsg.style.display = 'block';
        }
    });
});
