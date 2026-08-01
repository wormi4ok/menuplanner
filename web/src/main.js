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
import ConfigProgrammatic from 'buefy/dist/esm/config.js';
import Autocomplete from 'buefy/dist/esm/autocomplete.js';
import Button from 'buefy/dist/esm/button.js';
import Checkbox from 'buefy/dist/esm/checkbox.js';
import Collapse from 'buefy/dist/esm/collapse.js';
import Dialog from 'buefy/dist/esm/dialog.js';
import Field from 'buefy/dist/esm/field.js';
import Icon from 'buefy/dist/esm/icon.js';
import Input from 'buefy/dist/esm/input.js';
import Loading from 'buefy/dist/esm/loading.js';
import Modal from 'buefy/dist/esm/modal.js';
import Navbar from 'buefy/dist/esm/navbar.js';
import Notification from 'buefy/dist/esm/notification.js';
import Numberinput from 'buefy/dist/esm/numberinput.js';
import Radio from 'buefy/dist/esm/radio.js';
import Table from 'buefy/dist/esm/table.js';
import Tabs from 'buefy/dist/esm/tabs.js';
import Tag from 'buefy/dist/esm/tag.js';
import Toast from 'buefy/dist/esm/toast.js';
import Tooltip from 'buefy/dist/esm/tooltip.js';
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
ConfigProgrammatic.setOptions({
  defaultIconComponent: 'fa',
  defaultIconPack: 'fas',
});
[
  Autocomplete,
  Button,
  Checkbox,
  Collapse,
  Dialog,
  Field,
  Icon,
  Input,
  Loading,
  Modal,
  Navbar,
  Notification,
  Numberinput,
  Radio,
  Table,
  Tabs,
  Tag,
  Toast,
  Tooltip,
].forEach((plugin) => app.use(plugin));
app.use(createPinia());
app.use(router);

app.mount('#app');
