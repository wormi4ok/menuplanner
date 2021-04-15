<template>
  <div class="about">
    <h1>Recipes</h1>
    <div v-for="recipe in listRecipes" :key="recipe.id">
      <RecipeCard v-bind="recipe" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import RecipeCard from '@/components/RecipeCard.vue';

export default {
  name: 'Recipes',
  components: {
    RecipeCard,
  },
  computed: mapGetters(['listRecipes']),
};
</script>
