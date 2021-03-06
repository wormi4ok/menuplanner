import Vuex from 'vuex';
import Vue from 'vue';
import user from './modules/user';
import recipes from './modules/recipes';
import courses from './modules/courses';
import week from './modules/week';
import error from './modules/error';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    user,
    recipes,
    week,
    courses,
    error,
  },
});
