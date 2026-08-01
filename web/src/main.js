import { createApp } from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faCheck,
  faChevronDown,
  faEdit,
  faEnvelope,
  faLock,
  faPlusCircle,
  faSearchPlus,
  faTimes,
  faTimesCircle,
  faTrash,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import Vue3TouchEvents from 'vue3-touch-events';
import 'buefy/dist/css/buefy.css';
import './theme.css';
import Buefy from 'buefy';
import { createPinia } from 'pinia';
import App from './App.vue';
import { init as initGoogleAuth } from './auth/google';

import router from './router';

library.add(
  faCheck,
  faChevronDown,
  faEdit,
  faEnvelope,
  faLock,
  faPlusCircle,
  faSearchPlus,
  faTimes,
  faTimesCircle,
  faTrash,
);

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
app.use(createPinia());
app.use(router);

app.mount('#app');
