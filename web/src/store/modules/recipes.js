/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '@/api';

const state = () => ({
  recipes: [],
});

const courses = [
  'breakfast',
  'main',
  'pudding',
];

const getters = {
  listRecipes: (state) => state.recipes,
  listCourses: () => courses,
};

const actions = {
  async fetchRecipes({ commit }) {
    api.recipe.list().then((response) => {
      commit('setRecipes', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
  async createRecipe({ commit }, recipe) {
    api.recipe.create(recipe).then((response) => {
      commit('pushNewRecipe', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
};

const mutations = {
  setRecipes: (state, recipes) => {
    state.recipes = recipes;
  },
  pushNewRecipe: (state, recipe) => {
    state.recipes.push(recipe);
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
