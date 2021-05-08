<template>
  <div class="card grid-item">
    <header
      :class="{'card-header':true, 'is-clickable': description}"
      @click="toggleDescription"
      aria-controls="descriptionBlock">
      <p class="card-header-title">{{ name }}</p>
    </header>
    <div class="card-content">
      <section>
        <b-collapse
          aria-id="descriptionBlock"
          class="block"
          animation="slide"
          v-model="showDescription">
          {{ description }}
        </b-collapse>
        <b-taglist>
          <b-tag type="is-success">Prot: {{ totalProtein }}</b-tag>
          <b-tag type="is-info ">Fat: {{ totalFat }}</b-tag>
          <b-tag type="is-warning ">Carbs: {{ totalCarbs }}</b-tag>
          <b-tag type="is-link "> {{ totalCalories }} kcal</b-tag>
        </b-taglist>
      </section>
    </div>
    <footer class="card-footer">
      <a class="card-footer-item" @click="deleteRecipe">
        <b-icon icon="times"></b-icon>
      </a>
      <div class="card-footer-item">
        <b-tag>{{ recipeSize }}</b-tag>
      </div>
    </footer>
  </div>
</template>

<script>
import RecipeNutrition from '@/mixins/RecipeNutrition';

export default {
  name: 'Recipe',
  props: {
    id: Number,
    name: String,
    description: String,
    protein: Number,
    fat: Number,
    carbs: Number,
    calories: Number,
    portion: Number,
    quantity: Number,
  },
  data() {
    return {
      showDescription: false,
    };
  },
  computed: {
    recipeSize() {
      if (this.portion) {
        return `${this.portion} g`;
      }

      return '1 serving';
    },
    totalProtein() {
      return this.recipeProtein({
        protein: this.protein,
        quantity: this.quantity,
        portion: this.portion,
      });
    },
    totalFat() {
      return this.recipeFat({
        fat: this.fat,
        quantity: this.quantity,
        portion: this.portion,
      });
    },
    totalCarbs() {
      return this.recipeCarbs({
        carbs: this.carbs,
        quantity: this.quantity,
        portion: this.portion,
      });
    },
    totalCalories() {
      return this.recipeCalories({
        calories: this.calories,
        quantity: this.quantity,
        portion: this.portion,
      });
    },
  },
  methods: {
    deleteRecipe() {
      this.$emit('delete-recipe');
    },
    toggleDescription() {
      this.showDescription = !this.showDescription;
    },
  },
  mixins: [
    RecipeNutrition,
  ],
};
</script>

<style scoped>
.card {
  display: grid;
  grid-template-rows: min-content 1fr min-content;
}
</style>
