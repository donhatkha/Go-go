// API Configuration
const API_BASE_URL = window.location.origin + '/api';

// State
let users = [];

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    checkHealth();
    loadUsers();
    setupEventListeners();
});

// Event Listeners
function setupEventListeners() {
    // Add user form
    document.getElementById('user-form').addEventListener('submit', handleAddUser);

    // Refresh button
    document.getElementById('refresh-btn').addEventListener('click', loadUsers);

    // Edit modal
    const modal = document.getElementById('edit-modal');
    const closeBtn = document.querySelector('.close');
    
    closeBtn.addEventListener('click', () => {
        modal.style.display = 'none';
    });

    window.addEventListener('click', (event) => {
        if (event.target === modal) {
            modal.style.display = 'none';
        }
    });

    // Edit form
    document.getElementById('edit-form').addEventListener('submit', handleEditUser);
}

// API Functions
async function checkHealth() {
    const statusIndicator = document.querySelector('.status-indicator');
    const statusText = document.querySelector('.status-text');

    try {
        const response = await fetch(`${API_BASE_URL}/health`);
        const data = await response.json();

        if (data.success) {
            statusIndicator.classList.add('healthy');
            statusText.textContent = 'Server is healthy';
        } else {
            statusIndicator.classList.add('unhealthy');
            statusText.textContent = 'Server error';
        }
    } catch (error) {
        statusIndicator.classList.add('unhealthy');
        statusText.textContent = 'Server is offline';
        console.error('Health check failed:', error);
    }
}

async function loadUsers() {
    const container = document.getElementById('users-container');
    container.innerHTML = '<p class="loading">Loading users...</p>';

    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        const data = await response.json();

        if (data.success) {
            users = data.data || [];
            displayUsers(users);
        } else {
            container.innerHTML = `<p class="error">Failed to load users: ${data.message}</p>`;
        }
    } catch (error) {
        container.innerHTML = '<p class="error">Failed to connect to server</p>';
        console.error('Load users failed:', error);
    }
}

async function handleAddUser(e) {
    e.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;

    try {
        const response = await fetch(`${API_BASE_URL}/users`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email }),
        });

        const data = await response.json();

        if (data.success) {
            showNotification('User added successfully!', 'success');
            document.getElementById('user-form').reset();
            loadUsers();
        } else {
            showNotification(`Failed to add user: ${data.message}`, 'error');
        }
    } catch (error) {
        showNotification('Failed to connect to server', 'error');
        console.error('Add user failed:', error);
    }
}

async function handleEditUser(e) {
    e.preventDefault();

    const id = document.getElementById('edit-id').value;
    const name = document.getElementById('edit-name').value;
    const email = document.getElementById('edit-email').value;

    try {
        const response = await fetch(`${API_BASE_URL}/users/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email }),
        });

        const data = await response.json();

        if (data.success) {
            showNotification('User updated successfully!', 'success');
            document.getElementById('edit-modal').style.display = 'none';
            loadUsers();
        } else {
            showNotification(`Failed to update user: ${data.message}`, 'error');
        }
    } catch (error) {
        showNotification('Failed to connect to server', 'error');
        console.error('Update user failed:', error);
    }
}

async function deleteUser(id) {
    if (!confirm('Are you sure you want to delete this user?')) {
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/users/${id}`, {
            method: 'DELETE',
        });

        const data = await response.json();

        if (data.success) {
            showNotification('User deleted successfully!', 'success');
            loadUsers();
        } else {
            showNotification(`Failed to delete user: ${data.message}`, 'error');
        }
    } catch (error) {
        showNotification('Failed to connect to server', 'error');
        console.error('Delete user failed:', error);
    }
}

function editUser(user) {
    document.getElementById('edit-id').value = user.id;
    document.getElementById('edit-name').value = user.name;
    document.getElementById('edit-email').value = user.email;
    document.getElementById('edit-modal').style.display = 'block';
}

// Display Functions
function displayUsers(users) {
    const container = document.getElementById('users-container');

    if (users.length === 0) {
        container.innerHTML = '<p class="loading">No users found. Add a user to get started!</p>';
        return;
    }

    const usersHTML = users.map(user => {
        const date = new Date(user.created_at).toLocaleDateString('en-US', {
            year: 'numeric',
            month: 'short',
            day: 'numeric',
        });

        return `
            <div class="user-card">
                <h3>${escapeHtml(user.name)}</h3>
                <p><strong>Email:</strong> ${escapeHtml(user.email)}</p>
                <small><strong>ID:</strong> ${user.id}</small><br>
                <small><strong>Created:</strong> ${date}</small>
                <div class="user-actions">
                    <button class="btn btn-small btn-edit" data-user-id="${user.id}" data-user-name="${escapeHtml(user.name)}" data-user-email="${escapeHtml(user.email)}">
                        ✏️ Edit
                    </button>
                    <button class="btn btn-small btn-delete" data-user-id="${user.id}">
                        🗑️ Delete
                    </button>
                </div>
            </div>
        `;
    }).join('');

    container.innerHTML = `<div class="users-grid">${usersHTML}</div>`;

    // Add event listeners to all edit and delete buttons
    container.querySelectorAll('.btn-edit').forEach(button => {
        button.addEventListener('click', function() {
            const user = {
                id: parseInt(this.dataset.userId),
                name: this.dataset.userName,
                email: this.dataset.userEmail
            };
            editUser(user);
        });
    });

    container.querySelectorAll('.btn-delete').forEach(button => {
        button.addEventListener('click', function() {
            deleteUser(parseInt(this.dataset.userId));
        });
    });
}

function showNotification(message, type) {
    const container = document.querySelector('main');
    const notification = document.createElement('div');
    notification.className = type;
    notification.textContent = message;
    
    container.insertBefore(notification, container.firstChild);

    setTimeout(() => {
        notification.remove();
    }, 5000);
}

function escapeHtml(text) {
    const map = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#039;'
    };
    return text.replace(/[&<>"']/g, m => map[m]);
}
