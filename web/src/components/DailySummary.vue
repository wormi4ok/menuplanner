<template>
  <div class="block">
    <ul>
      <li>Protein: {{ totalProtein }}</li>
      <li>Fat: {{ totalFat }}</li>
      <li>Carbs: {{ totalCarbs }}</li>
      <li>
        <span :class="isOverLimit ? 'has-text-danger' : ''">Calories:{{ totalCalories }}</span>
      </li>
    </ul>
  </div>
</template>

<script>
import RecipeNutrition from '@/mixins/RecipeNutrition';

export default {
  name: 'DailySummary',
  props: {
    recipes: Object,
  },
  computed: {
    totalProtein() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + this.recipeProtein(r) : t), 0);
    },
    totalFat() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + this.recipeFat(r) : t), 0);
    },
    totalCarbs() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + this.recipeCarbs(r) : t), 0);
    },
    totalCalories() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + this.recipeCalories(r) : t), 0);
    },
    isOverLimit() {
      return this.totalCalories > 1700;
    },
  },
  mixins: [
    RecipeNutrition,
  ],
};
</script>

<style scoped>

</style>
