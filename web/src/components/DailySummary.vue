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
import RecipeCalories from '@/mixins/RecipeCalories';

export default {
  name: 'DailySummary',
  props: {
    recipes: Object,
  },
  computed: {
    totalProtein() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + r.protein : t), 0);
    },
    totalFat() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + r.fat : t), 0);
    },
    totalCarbs() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + r.carbs : t), 0);
    },
    totalCalories() {
      return Object.values(this.recipes).reduce((t, r) => (r ? t + this.recipeCalories(r) : t), 0);
    },
    isOverLimit() {
      return this.totalCalories > 1700;
    },
  },
  mixins: [
    RecipeCalories,
  ],
};
</script>

<style scoped>

</style>
