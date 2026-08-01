const GIS_URL = 'https://accounts.google.com/gsi/client';

let codeClient = null;
let pending = null;

function settle(outcome, value) {
  if (!pending) return;
  const handler = pending[outcome];
  pending = null;
  handler(value);
}

function loadScript() {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script');
    script.src = GIS_URL;
    script.async = true;
    script.onload = resolve;
    script.onerror = () => reject(new Error('Failed to load Google Identity Services'));
    document.head.appendChild(script);
  });
}

export async function init(clientId) {
  try {
    await loadScript();
    codeClient = window.google.accounts.oauth2.initCodeClient({
      client_id: clientId,
      scope: 'openid profile email',
      ux_mode: 'popup',
      select_account: true,
      callback: (response) => {
        if (response.error) {
          settle('reject', new Error(response.error_description || response.error));
        } else {
          settle('resolve', response.code);
        }
      },
      error_callback: (error) => {
        settle('reject', new Error(error.message || 'Google Sign-In did not complete'));
      },
    });
  } catch (e) {
    console.error(e);
  }
}

export function getAuthCode() {
  if (!codeClient) return Promise.reject(new Error('Google Sign-In is not ready yet'));

  return new Promise((resolve, reject) => {
    settle('reject', new Error('Superseded by a new Google Sign-In request'));
    pending = { resolve, reject };
    codeClient.requestCode();
  });
}
