<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ name }}</p>
      <p class="card-header-icon" aria-label="delete">
        <b-button icon-right="times" type="is-ghost is-danger" inverted @click="deleteRecipe"/>
      </p>
    </header>
    <div class="card-content">
      <section>
        <b-progress :max="total" show-value>
          <template #bar>
            <b-progress-bar type="is-success" :value="protein" show-value>P</b-progress-bar>
            <b-progress-bar type="is-info" :value="fat">F</b-progress-bar>
            <b-progress-bar type="is-warning" :value="carbs" show-value>C</b-progress-bar>
          </template>
        </b-progress>
        <b-progress :value="calories" :max="1700" show-value>Calories</b-progress>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Recipe',
  props: {
    id: Number,
    name: String,
    protein: Number,
    fat: Number,
    carbs: Number,
    calories: Number,
  },
  computed: {
    total() {
      return this.protein + this.fat + this.carbs;
    },
  },
  methods: {
    deleteRecipe() {
      this.$emit('delete-recipe');
    },
  },
};
</script>

<style scoped>

</style>
