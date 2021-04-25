<template>
  <div class="menu">
    <MenuSlot v-for="(recipe, slot) in recipes"
              :key="slot"
              :recipe="recipe"
              :course="course(slot)"
              @delete-recipe="deleteRecipe(slot)"
              @pick-recipe="pickRecipe(slot, $event)"/>
    <div class="block">
      Summary
    </div>
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'DailyMenu',
  components: {
    MenuSlot,
  },
  props: {
    recipes: Object,
  },
  computed: {
    ...mapGetters([
      'listCourses',
    ]),
  },
  methods: {
    deleteRecipe(slot) {
      this.$emit('empty-slot', slot);
    },
    pickRecipe(slot, recipe) {
      this.$emit('fill-slot', { slot, recipe });
    },
    course(slot) {
      const map = {
        0: 'breakfast',
        1: 'main',
        2: 'main',
        3: 'pudding',
      };
      return this.listCourses.find((course) => course.name === map[slot]);
    },
  },
};
</script>

<style scoped>

</style>
