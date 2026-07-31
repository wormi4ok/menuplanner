import { createStore } from 'vuex';
import user from './modules/user';
import recipes from './modules/recipes';
import courses from './modules/courses';
import week from './modules/week';
import error from './modules/error';

export default createStore({
  modules: {
    user,
    recipes,
    week,
    courses,
    error,
  },
});
