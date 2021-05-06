<template>
  <RecipeCard v-if="!isEmpty" v-bind="recipe" @delete-recipe="deleteRecipe"/>
  <EmptySlot v-else-if="course" :course="course" @pick-recipe="pickRecipe($event)"/>
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
    course: Object,
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
