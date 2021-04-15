/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import recipeApi from '@/api';

const state = () => ({
  recipes: [],
});

const getters = {
  listRecipes: (state) => state.recipes,
};

const actions = {
  async fetchRecipes({ commit }) {
    const response = await recipeApi.list();
    commit('setRecipes', response.data);
  },
};

const mutations = {
  setRecipes: (state, recipes) => {
    state.recipes = recipes;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
