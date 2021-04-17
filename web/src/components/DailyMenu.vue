<template>
  <div class="menu">
    <div class="block">
      <RecipeCard v-if="recipes[0]" v-bind="recipes[0]" @delete-recipe="deleteRecipe(0)"/>
      <EmptySlot v-else @pick-recipe="pickRecipe(0, $event)"></EmptySlot>
    </div>
    <div class="block">
      <RecipeCard v-if="recipes[1]" v-bind="recipes[1]" @delete-recipe="deleteRecipe(1)"/>
      <EmptySlot v-else @pick-recipe="pickRecipe(1, $event)"></EmptySlot>
    </div>
    <div class="block">
      <RecipeCard v-if="recipes[2]" v-bind="recipes[2]" @delete-recipe="deleteRecipe(2)"/>
      <EmptySlot v-else @pick-recipe="pickRecipe(2, $event)"></EmptySlot>
    </div>
    <div class="block">
      Summary
    </div>
  </div>
</template>

<script>
import RecipeCard from '@/components/RecipeCard.vue';
import EmptySlot from '@/components/EmptySlot.vue';

export default {
  name: 'Weekday',
  components: {
    RecipeCard,
    EmptySlot,
  },
  props: {
    recipes: Object,
  },
  methods: {
    deleteRecipe(slot) {
      this.$emit('empty-slot', slot);
    },
    pickRecipe(slot, recipe) {
      this.$emit('fill-slot', { slot, recipe });
    },
  },
};
</script>

<style scoped>

</style>
