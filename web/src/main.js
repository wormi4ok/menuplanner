import Vue from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import 'buefy/dist/buefy.css';
import Buefy from 'buefy';
import App from './App.vue';
import store from './store';

Vue.component('fa', FontAwesomeIcon);
library.add(fas);

Vue.use(Buefy, {
  defaultIconComponent: 'fa',
  defaultIconPack: 'fas',
});

Vue.config.productionTip = false;

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');
