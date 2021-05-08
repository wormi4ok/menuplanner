export default {
  methods: {
    recipeProtein(recipe) {
      return this.recipeAmount(recipe, 'protein');
    },
    recipeFat(recipe) {
      return this.recipeAmount(recipe, 'fat');
    },
    recipeCarbs(recipe) {
      return this.recipeAmount(recipe, 'carbs');
    },
    recipeCalories(recipe) {
      return this.recipeAmount(recipe, 'calories');
    },
    recipeAmount: (recipe, property) => {
      if (!recipe || !recipe[property]) {
        return 0;
      }

      if (recipe.quantity) {
        return Math.ceil((recipe[property] / recipe.quantity) * recipe.portion);
      }

      return recipe[property];
    },
  },
};
