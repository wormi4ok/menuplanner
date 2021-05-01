export default {
  methods: {
    recipeCalories: (recipe) => {
      if (!recipe) {
        return 0;
      }

      if (recipe.quantity) {
        return Math.ceil((recipe.calories / recipe.quantity) * recipe.portion);
      }

      return recipe.calories;
    },
  },
};
