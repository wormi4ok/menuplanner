import { createApp } from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import Vue3TouchEvents from 'vue3-touch-events';
import 'buefy/dist/css/buefy.css';
import './theme.css';
import Buefy from 'buefy';
import App from './App.vue';
import store from './store';
import { init as initGoogleAuth } from './auth/google';

import router from './router';

library.add(fas);

if (window.config.MP_CLIENT_ID) {
  initGoogleAuth(window.config.MP_CLIENT_ID);
}

const app = createApp(App);

app.component('fa', FontAwesomeIcon);
app.use(Vue3TouchEvents, {});
app.use(Buefy, {
  defaultIconComponent: 'fa',
  defaultIconPack: 'fas',
});
app.use(store);
app.use(router);

app.mount('#app');
