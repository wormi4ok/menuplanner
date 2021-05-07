<template>
  <div class="card">
    <b-button
      v-if="!showInput"
      type="is-success"
      icon-right="plus-circle"
      size="is-large" inverted
      @click="onShowRecipeInput"
    />
    <div v-show="showInput" class="has-text-centered">
      <b-field label="Pick a recipe">
        <b-autocomplete
          v-model="name"
          :data="filteredDataObj"
          ref="recipeInput"
          type="search"
          icon="search-plus"
          field="name"
          placeholder="Recipe name"
          open-on-focus rounded clearable
          @select="pickRecipe">
          <template #empty>No recipes found</template>
        </b-autocomplete>
      </b-field>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'EmptySlot',
  props: {
    course: {
      type: Object,
      required: true,
    },
  },
  data: () => ({
    name: '',
    showInput: false,
  }),
  computed: {
    filteredDataObj() {
      return this.recipes.filter((option) => (
        option.name
          .toString()
          .toLowerCase()
          .indexOf(this.name.toLowerCase()) >= 0
      ));
    },
    recipes() {
      return this.recipesByCourse(this.course);
    },
    ...mapGetters([
      'recipesByCourse',
    ]),
  },
  methods: {
    pickRecipe(option) {
      this.$emit('pick-recipe', option);
    },
    onShowRecipeInput() {
      this.showInput = true;
      this.$refs.recipeInput.focus();
    },
  },
};
</script>

<style scoped>
.card {
  padding: 0.75rem 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
