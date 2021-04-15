import Vuex from 'vuex';
import Vue from 'vue';
import recipes from './modules/recipes';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    recipes,
  },
});
