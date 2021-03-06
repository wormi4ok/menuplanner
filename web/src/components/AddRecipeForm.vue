<template>
  <div class="modal-card" style="width: 680px">
    <form @submit.prevent="submitRecipeForm">
      <header class="modal-card-head">
        <p class="modal-card-title">{{ isUpdate ? 'Update recipe' : 'Add new' }}</p>
        <button type="button" class="delete" @click="$emit('close')"/>
      </header>
      <section class="modal-card-body">
        <b-notification
          v-if="error"
          type="is-danger is-light"
          aria-close-label="Close error message"
          role="alert">
          {{ errorMsg }}
        </b-notification>
        <b-field>
          <b-field label="Recipe" expanded>
            <b-input
              v-model="name"
              validation-message="Recipe name is very important!"
              type="text"
              required
              expanded/>
            <b-checkbox-button v-model="courses"
                               v-for="course in listCourses"
                               :value="course.name"
                               :key="course.id"
                               :native-value="course"
                               :type="courseColorCode(course)">
              <span class="is-capitalized">{{ course.name }}</span>
            </b-checkbox-button>
          </b-field>
        </b-field>
        <b-field label="Description">
          <b-input type="textarea" v-model="description"></b-input>
        </b-field>
        <b-field label="Calories">
          <b-numberinput v-model="calories" step="50" min-step="1" min="1" message="Calories"/>
        </b-field>
        <b-field grouped>
          <b-field label="Protein">
            <b-numberinput
              v-model="protein"
              step="5"
              min-step="1"
              controls-position="compact"
              min="1"
              size="is-small"/>
          </b-field>
          <b-field label="Fat">
            <b-numberinput v-model="fat"
                           step="5"
                           min-step="1"
                           min="1"
                           controls-position="compact"
                           size="is-small"/>
          </b-field>
          <b-field label="Carbs">
            <b-numberinput v-model="carbs"
                           step="5"
                           min-step="1"
                           min="1"
                           controls-position="compact"
                           size="is-small"/>
          </b-field>
        </b-field>
        <b-field aria-controls="quantityPerServing" position="is-centered">
          <b-radio-button
            v-model="showQuantityInput"
            :native-value="false"
            type="is-primary is-light is-outlined">
            <span>Per serving</span>
          </b-radio-button>
          <b-radio-button
            v-model="showQuantityInput"
            :native-value="true"
            type="is-info is-light is-outlined">
            <span>Specify Quantity</span>
          </b-radio-button>
        </b-field>
        <b-collapse aria-id="quantityPerServing" class="panel" v-model="showQuantityInput">
          <b-field grouped position="is-centered">
            <b-field label="Quantity (g)">
              <b-numberinput v-model="quantity"
                             step="50"
                             min-step="1"
                             controls-position="compact"/>
            </b-field>
            <b-field label="Portion (g)">
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
        <b-button :label="isUpdate ? 'Update' : 'Create'" class="is-primary" native-type="submit"/>
      </footer>
    </form>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import CourseColor from '@/mixins/CourseColor';

export default {
  name: 'AddRecipeForm',
  props: {
    recipe: Object,
  },
  data: () => ({
    id: 0,
    name: '',
    courses: [],
    description: '',
    imageUrl: '',
    calories: 0,
    protein: 0,
    fat: 0,
    carbs: 0,
    quantity: 100,
    portion: 350,
    showQuantityInput: false,
    error: '',
  }),
  created() {
    if (this.recipe) {
      this.id = this.recipe.id;
      this.name = this.recipe.name;
      this.recipe.courses.forEach((course, index) => {
        this.$set(this.courses, index, course);
      });

      this.description = this.recipe.description;
      this.imageUrl = this.recipe.imageUrl;
      this.calories = this.recipe.calories;
      this.protein = this.recipe.protein;
      this.fat = this.recipe.fat;
      this.carbs = this.recipe.carbs;

      this.quantity = this.recipe.quantity;
      this.portion = this.recipe.portion;
      if (this.quantity || this.portion) {
        this.showQuantityInput = true;
      }
    }
  },
  computed: {
    ...mapGetters([
      'listCourses',
    ]),
    isUpdate() {
      return this.id > 0;
    },
    errorMsg() {
      let msg = '';
      if (Object.prototype.hasOwnProperty.call(this.error, 'errors')) {
        this.error.errors.forEach((error) => {
          msg += `${error.message}: ${error.field}\n`;
        });
      } else {
        msg = this.error;
      }
      return msg;
    },
  },
  methods: {
    ...mapActions([
      'createRecipe',
      'updateRecipe',
    ]),
    async submitRecipeForm() {
      this.error = '';
      const recipe = {
        name: this.name,
        courses: this.courses,
        description: this.description,
        imageUrl: this.imageUrl,
        calories: this.calories,
        protein: this.protein,
        fat: this.fat,
        carbs: this.carbs,
        quantity: this.showQuantityInput ? this.quantity : 0,
        portion: this.showQuantityInput ? this.portion : 0,
      };
      try {
        if (this.id > 0) {
          recipe.id = this.id;
          await this.updateRecipe(recipe);
        } else {
          await this.createRecipe(recipe);
        }
        this.$emit('close');
      } catch (e) {
        this.error = e.response.data;
      }
    },
  },
  mixins: [
    CourseColor,
  ],
};
</script>

<style scoped>

</style>
