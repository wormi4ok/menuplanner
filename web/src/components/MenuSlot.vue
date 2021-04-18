<template>
  <div class="block">
    <RecipeCard v-if="recipe" v-bind="recipe" @delete-recipe="deleteRecipe"/>
    <EmptySlot v-else @pick-recipe="pickRecipe($event)"></EmptySlot>
  </div>
</template>

<script>
import RecipeCard from '@/components/RecipeCard.vue';
import EmptySlot from '@/components/EmptySlot.vue';

export default {
  name: 'MenuSlot',
  components: {
    RecipeCard,
    EmptySlot,
  },
  props: {
    recipe: Object,
  },
  methods: {
    deleteRecipe() {
      this.$emit('delete-recipe');
    },
    pickRecipe(recipe) {
      this.$emit('pick-recipe', recipe);
    },
  },
};
</script>

<style scoped>

</style>
