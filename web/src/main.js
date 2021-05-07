import Vue from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import Vue2TouchEvents from 'vue2-touch-events';
import 'buefy/dist/buefy.css';
import Buefy from 'buefy';
import App from './App.vue';
import store from './store';

import router from './router';

Vue.component('fa', FontAwesomeIcon);
library.add(fas);
Vue.use(Vue2TouchEvents);
Vue.use(Buefy, {
  defaultIconComponent: 'fa',
  defaultIconPack: 'fas',
});

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  data() {
    return {
      isMobile: window.innerWidth <= 768,
    };
  },
  render: (h) => h(App),
}).$mount('#app');
