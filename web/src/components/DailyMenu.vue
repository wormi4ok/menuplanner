<template>
  <div class="menu">
    <MenuSlot v-for="(recipe, slot) in recipes"
              :key="slot"
              :recipe="recipe"
              :course="course(slot)"
              @delete-recipe="deleteRecipe(slot)"
              @pick-recipe="pickRecipe(slot, $event)"/>
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
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'DailyMenu',
  components: {
    MenuSlot,
  },
  props: {
    recipes: Object,
  },
  computed: {
    ...mapGetters([
      'listCourses',
    ]),
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
      return Object.values(this.recipes).reduce((t, r) => (r ? t + r.calories : t), 0);
    },
    isOverLimit() {
      return this.totalCalories > 1700;
    },
  },
  methods: {
    deleteRecipe(slot) {
      this.$emit('empty-slot', slot);
    },
    pickRecipe(slot, recipe) {
      this.$emit('fill-slot', { slot, recipe });
    },
    course(slot) {
      const map = {
        0: 'breakfast',
        1: 'main',
        2: 'main',
        3: 'pudding',
      };
      return this.listCourses.find((course) => course.name === map[slot]);
    },
  },
};
</script>

<style scoped>

</style>
