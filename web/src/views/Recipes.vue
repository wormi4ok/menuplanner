<template>
  <b-table
    :data="data"
    detailed
    detail-key="id"
    :show-detail-icon="false"
    striped>

    <b-table-column field="name" label="Recipe" width="80" v-slot="props">
      <a @click="props.toggleDetails(props.row)">
        {{ props.row.name }}
      </a>
    </b-table-column>

    <b-table-column field="courses" label="Course" width="30" v-slot="props">
      <b-tag
        v-for="course in props.row.courses"
        :key="course.id"
        :type="courseColorCode(course)"
        rounded>
        {{ course.name }}
      </b-tag>
    </b-table-column>

    <b-table-column field="calories" label="Calories" width="40" numeric v-slot="props">
      {{ props.row.calories }}
    </b-table-column>

    <b-table-column field="protein" label="Protein" width="40" numeric v-slot="props">
      {{ props.row.protein }}
    </b-table-column>

    <b-table-column field="fat" label="Fat" width="40" numeric v-slot="props">
      {{ props.row.fat }}
    </b-table-column>

    <b-table-column field="carbs" label="Carbs" width="40" numeric v-slot="props">
      {{ props.row.carbs }}
    </b-table-column>

    <b-table-column field="quantity" label="Quantity" width="40" numeric v-slot="props">
      {{ props.row.quantity }}
    </b-table-column>

    <b-table-column field="portion" label="Portion" width="40" numeric v-slot="props">
      {{ props.row.portion }}
    </b-table-column>

    <b-table-column field="id" label="Actions" width="20" v-slot="props">
      <b-button type="is-ghost" icon-right="edit" @click="onEdit(props.row)"/>
      <b-button type="is-danger" inverted icon-right="trash" @click="onDelete(props.row)"/>
    </b-table-column>

    <template #detail="props">
      {{ props.row.description }}
    </template>

    <template #empty>
      <div class="has-text-centered">No recipes</div>
    </template>
  </b-table>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import CourseColor from '@/mixins/CourseColor';
import AddRecipeForm from '@/components/AddRecipeForm.vue';

export default {
  name: 'Recipes',
  computed: mapGetters({
    data: 'listRecipes',
    recipeInUse: 'recipePosition',
  }),
  created() {
    if (this.data.length === 0) {
      this.fetchRecipes();
    }
  },
  methods: {
    ...mapActions([
      'fetchRecipes',
      'deleteRecipe',
      'emptySlot',
    ]),
    onDelete(recipe) {
      const positions = this.recipeInUse(recipe.id);
      if (positions) {
        this.$buefy.dialog.confirm({
          message: `The recipe <b>${recipe.name}</b> is used in the current week. Remove anyways?`,
          onConfirm: () => {
            positions.map((position) => this.emptySlot(position));
            this.deleteRecipe(recipe.id);
          },
        });
      } else {
        this.$buefy.dialog.confirm({
          message: `Remove <b>${recipe.name}</b>?`,
          onConfirm: () => this.deleteRecipe(recipe.id),
        });
      }
    },
    onEdit(recipe) {
      this.$buefy.modal.open({
        parent: this,
        component: AddRecipeForm,
        props: {
          recipe,
        },
        hasModalCard: true,
      });
    },
  },
  mixins: [
    CourseColor,
  ],
};
</script>
