<template>
  <div class="menu">
    <div class="block">
      <RecipeCard v-if="recipes[0]" v-bind="recipes[0]"/>
      <EmptySlot v-else></EmptySlot>
    </div>
    <div class="block">
      <RecipeCard v-if="recipes[1]" v-bind="recipes[1]"/>
      <EmptySlot v-else></EmptySlot>
    </div>
    <div class="block">
      <RecipeCard v-if="recipes[2]" v-bind="recipes[2]"/>
      <EmptySlot v-else></EmptySlot>
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
};
</script>

<style scoped>

</style>
