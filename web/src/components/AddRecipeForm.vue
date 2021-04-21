<template>
  <div class="modal-card" style="width: 680px">
    <form @submit.prevent="addRecipe">
      <header class="modal-card-head">
        <p class="modal-card-title">Add new</p>
        <button type="button" class="delete" @click="$emit('close')"/>
      </header>
      <section class="modal-card-body">
        <b-field>
          <b-field label="Recipe" expanded>
            <b-input
              v-model="name"
              validation-message="Recipe name is very important!"
              type="text"
              required
              expanded/>
            <b-checkbox-button v-model="courses"
                               v-for="value in listCourses"
                               :value="value"
                               :key="value"
                               :native-value="value"
                               :type="courseColorCode(value)">
              <span class="is-capitalized">{{ value }}</span>
            </b-checkbox-button>
          </b-field>
        </b-field>
        <b-field label="Description">
          <b-input type="textarea" v-model="description"></b-input>
        </b-field>
        <b-field label="Calories">
          <b-numberinput v-model="calories" step="50" min-step="1" message="Calories"/>
        </b-field>
        <b-field grouped>
          <b-field label="Protein">
            <b-numberinput
              v-model="protein"
              step="50"
              min-step="1"
              controls-position="compact"
              size="is-small"/>
          </b-field>
          <b-field label="Fat">
            <b-numberinput v-model="fat"
                           step="50"
                           min-step="1"
                           controls-position="compact"
                           size="is-small"/>
          </b-field>
          <b-field label="Carbs">
            <b-numberinput v-model="carbs"
                           step="50"
                           min-step="1"
                           controls-position="compact"
                           size="is-small"/>
          </b-field>
        </b-field>
        <b-collapse :open="false" position="is-bottom" aria-id="contentIdForA11y1">
          <template #trigger="data">
            <a aria-controls="contentIdForA11y1">
              <b-icon :icon="!data.showQuantityInput ? 'menu-down' : 'menu-up'"></b-icon>
              Specify Quantity
            </a>
          </template>
          <b-field grouped position="is-centered">
            <b-field label="Quantity">
              <b-numberinput v-model="quantity"
                             step="50"
                             min-step="1"
                             controls-position="compact"/>
            </b-field>
            <b-field label="Portion">
              <b-numberinput v-model="portion"
                             step="50"
                             min-step="1"
                             controls-position="compact"/>
            </b-field>
          </b-field>
        </b-collapse>
      </section>
      <footer class="modal-card-foot">
        <b-button label="Close" @click="$emit('close')"/>
        <b-button label="Create" class="is-primary" native-type="submit"/>
      </footer>
    </form>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'AddRecipeForm',
  data: () => ({
    name: '',
    courses: [],
    description: '',
    imageUrl: '',
    calories: 0,
    protein: 0,
    fat: 0,
    carbs: 0,
    quantity: 0,
    portion: 0,
    showQuantityInput: false,
  }),
  computed: {
    ...mapGetters([
      'listCourses',
    ]),
  },
  methods: {
    ...mapActions([
      'createRecipe',
    ]),
    addRecipe() {
      const recipe = {
        name: this.name,
        courses: this.courses,
        description: this.description,
        imageUrl: this.imageUrl,
        calories: this.calories,
        protein: this.protein,
        fat: this.fat,
        carbs: this.carbs,
        quantity: this.quantity,
        portion: this.portion,
      };
      this.createRecipe(recipe);
      this.$emit('close');
    },
    courseColorCode: (course) => {
      const map = {
        breakfast: 'is-warning',
        main: 'is-info',
        pudding: 'is-danger',
      };
      return map[course];
    },
  },
};
</script>

<style scoped>

</style>
