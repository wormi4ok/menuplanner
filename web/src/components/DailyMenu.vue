<template>
  <div class="menu">
    <MenuSlot v-for="(recipe, slot) in recipes"
              :key="slot"
              :recipe="recipe"
              @delete-recipe="deleteRecipe(slot)"
              @pick-recipe="pickRecipe(slot, $event)"/>
    <div class="block">
      Summary
    </div>
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';

export default {
  name: 'DailyMenu',
  components: {
    MenuSlot,
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
