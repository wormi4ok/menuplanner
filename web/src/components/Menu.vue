<template>
  <div class="menu">
    <div class="block">
      <Recipe v-bind="recipes[0]"/>
    </div>
    <div class="block">
      <Recipe v-bind="recipes[1]"/>
    </div>
    <div class="block">
      <Recipe v-bind="recipes[2]"/>
    </div>
    <div class="block">
      Summary
    </div>
  </div>
</template>

<script>
import Recipe from '@/components/Recipe.vue';

export default {
  name: 'Weekday',
  components: {
    Recipe,
  },
  props: {
    recipes: Object,
  },
};
</script>

<style scoped>

</style>
