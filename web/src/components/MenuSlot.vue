<template>
  <div class="block">
    <RecipeCard v-if="!isEmpty" v-bind="recipe" @delete-recipe="deleteRecipe"/>
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
  computed: {
    isEmpty() {
      return !this.recipe || this.recipe.id === 0;
    },
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
