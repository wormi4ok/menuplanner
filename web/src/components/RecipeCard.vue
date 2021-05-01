<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ name }}</p>
    </header>
    <div class="card-content">
      <section>

        <b-taglist attached>
          <b-tag type="is-success">Protein: {{ protein }}</b-tag>
          <b-tag type="is-info">Fat: {{ fat }}</b-tag>
          <b-tag type="is-warning">Carbs: {{ carbs }}</b-tag>
          <b-tag>{{ recipeSize }}</b-tag>
        </b-taglist>
      </section>
    </div>
    <footer class="card-footer">
      <a class="card-footer-item" @click="deleteRecipe">
        <b-icon icon="times"></b-icon>
      </a>
    </footer>
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
    portion: Number,
    quantity: Number,
  },
  computed: {
    recipeSize() {
      if (this.portion) {
        return `${this.portion} g`;
      }

      return '1 serving';
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
