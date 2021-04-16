import Vuex from 'vuex';
import Vue from 'vue';
import recipes from './modules/recipes';
import week from './modules/week';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    recipes,
    week,
  },
});
