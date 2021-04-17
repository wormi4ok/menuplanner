<template>
  <div class="about">
    <h1>Recipes</h1>
    <div v-for="recipe in data" :key="recipe.id">
      <RecipeCard v-bind="recipe"/>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import RecipeCard from '@/components/RecipeCard.vue';

export default {
  name: 'Recipes',
  components: {
    RecipeCard,
  },
  computed: mapGetters({
    data: 'listRecipes',
  }),
  created() {
    if (this.data.length === 0) {
      this.fetchRecipes();
    }
  },
  methods: {
    ...mapActions(['fetchRecipes']),
  },
};
</script>
