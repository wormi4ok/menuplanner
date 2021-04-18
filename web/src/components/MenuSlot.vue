<template>
  <div class="block">
    <RecipeCard v-if="watchRecipe" v-bind="watchRecipe" @delete-recipe="deleteRecipe"/>
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
  data() {
    return {
      watchRecipe: this.recipe,
    };
  },
  methods: {
    deleteRecipe() {
      this.watchRecipe = null;
      this.$emit('delete-recipe');
    },
    pickRecipe(recipe) {
      this.watchRecipe = recipe;
      this.$emit('pick-recipe', recipe);
    },
  },
};
</script>

<style scoped>

</style>
