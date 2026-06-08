// Requires auth.js and profile.js to be loaded first

document.addEventListener('DOMContentLoaded', async () => {
    // Only run if we are on the dashboard
    const dashboardContainer = document.querySelector('.dashboard-container');
    if (!dashboardContainer) return;

    requireAuth();

    const userNameSpan = document.getElementById('userName');
    const roleBadge = document.getElementById('roleBadge');
    const logoutBtn = document.getElementById('logoutBtn');
    const adminSection = document.getElementById('adminSection');
    const viewUsersBtn = document.getElementById('viewUsersBtn');
    const usersList = document.getElementById('usersList');

    // Logout
    logoutBtn.addEventListener('click', logout);

    try {
        const user = await getProfile();
        userNameSpan.textContent = user.name;
        
        roleBadge.textContent = user.role;
        roleBadge.className = `role-badge role-${user.role}`;

        if (user.role === 'ADMIN') {
            adminSection.style.display = 'block';
        }
    } catch (err) {
        console.error(err);
    }

    if (viewUsersBtn) {
        viewUsersBtn.addEventListener('click', async () => {
            try {
                const res = await apiFetch('/admin/users');
                const data = await res.json();
                if (res.ok) {
                    let html = '<ul style="list-style-type: none; padding: 0;">';
                    data.users.forEach(u => {
                        html += `<li style="padding: 0.5rem; border-bottom: 1px solid #ddd;">
                            <strong>${u.name}</strong> (${u.email}) - <span class="role-badge role-${u.role}" style="font-size: 0.7em;">${u.role}</span>
                        </li>`;
                    });
                    html += '</ul>';
                    usersList.innerHTML = html;
                } else {
                    usersList.innerHTML = `<p style="color:red;">Error: ${data.error}</p>`;
                }
            } catch (err) {
                usersList.innerHTML = `<p style="color:red;">Error fetching users.</p>`;
            }
        });
    }
});
