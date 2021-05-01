/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '@/api';

const state = () => ({
  recipes: [],
});

const getters = {
  listRecipes: (state) => state.recipes,
  recipesByCourse: (state) => (course) => state.recipes.filter(
    (recipe) => recipe.courses.some(
      (c) => c.id === course.id,
    ),
  ),
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
  async updateRecipe({ commit }, recipe) {
    api.recipe.update(recipe.id, recipe).then((response) => {
      commit('modifyRecipe', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
  async deleteRecipe({ commit }, id) {
    api.recipe.delete(id).then(() => {
      commit('removeRecipe', id);
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
  modifyRecipe: (state, recipe) => {
    state.recipes = state.recipes.map((current) => {
      if (current.id === recipe.id) {
        return recipe;
      }

      return current;
    });
  },
  removeRecipe: (state, id) => {
    state.recipes = state.recipes.filter((recipes) => recipes.id !== id);
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
