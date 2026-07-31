const API_URL = 'https://apis.google.com/js/api.js';

let authInstance = null;

function loadScript() {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script');
    script.src = API_URL;
    script.onload = resolve;
    script.onerror = () => reject(new Error('Failed to load Google API script'));
    document.head.appendChild(script);
  });
}

function initClient(clientId) {
  return new Promise((resolve, reject) => {
    window.gapi.load('auth2', () => {
      window.gapi.auth2
        .init({ clientId, scope: 'profile email', prompt: 'select_account' })
        .then(resolve, reject);
    });
  });
}

export async function init(clientId) {
  try {
    await loadScript();
    await initClient(clientId);
    authInstance = window.gapi.auth2.getAuthInstance();
  } catch (e) {
    console.error(e);
  }
}

export async function getAuthCode() {
  if (!authInstance) throw new Error('Google Sign-In is not ready yet');
  const { code } = await authInstance.grantOfflineAccess({ prompt: 'select_account' });
  return code;
}
