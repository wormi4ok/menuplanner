import { defineStore } from 'pinia';
import api from '@/api';
import { useErrorStore } from './error';

export const useRecipesStore = defineStore('recipes', {
  state: () => ({
    recipes: [],
  }),
  getters: {
    listRecipes: (state) => state.recipes,
    recipesByCourse: (state) => (course) => state.recipes.filter(
      (recipe) => recipe.courses && recipe.courses.some(
        (c) => c.id === course.id,
      ),
    ),
  },
  actions: {
    async fetchRecipes() {
      try {
        const response = await api.recipe.list();
        this.recipes = response.data;
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
    async createRecipe(recipe) {
      const response = await api.recipe.create(recipe);
      this.recipes.push(response.data);
    },
    async updateRecipe(recipe) {
      const response = await api.recipe.update(recipe.id, recipe);
      this.recipes = this.recipes.map(
        (current) => (current.id === response.data.id ? response.data : current),
      );
    },
    async deleteRecipe(id) {
      try {
        await api.recipe.delete(id);
        this.recipes = this.recipes.filter((recipe) => recipe.id !== id);
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
  },
});

export default useRecipesStore;
