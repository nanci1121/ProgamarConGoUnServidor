document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('registerForm');
    const loginForm = document.getElementById('loginForm');
    const messageDiv = document.getElementById('message');

    const API_BASE_URL = 'https://localhost:8443';

    const showMessage = (message, type) => {
        messageDiv.textContent = message;
        messageDiv.className = `message ${type}`;
    };

    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const formData = new FormData(registerForm);
            const data = new URLSearchParams(formData);

            try {
                const response = await fetch(`${API_BASE_URL}/register`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                    body: data,
                });

                const responseText = await response.text();

                if (response.ok) {
                    showMessage(responseText, 'success');
                    registerForm.reset();
                } else {
                    showMessage(`Error: ${responseText}`, 'error');
                }
            } catch (error) {
                showMessage(`Error de conexión: ${error.message}`, 'error');
            }
        });
    }

    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const formData = new FormData(loginForm);
            const data = new URLSearchParams(formData);

            try {
                const response = await fetch(`${API_BASE_URL}/login`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                    body: data,
                });

                const responseText = await response.text();

                if (response.ok) {
                    showMessage(responseText, 'success');
                    loginForm.reset();
                } else {
                    showMessage(`Error: ${responseText}`, 'error');
                }
            } catch (error) {
                showMessage(`Error de conexión: ${error.message}`, 'error');
            }
        });
    }
});
