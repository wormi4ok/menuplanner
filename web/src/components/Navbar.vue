<template>
  <b-navbar class="mb-2">
    <template #brand>

      <b-navbar-item tag="router-link" :to="{ name: 'Home' }">
        <img src="@/assets/flaningo.jpeg" alt="MenuPlanner logo"/>
      </b-navbar-item>

    </template>
    <template #start>

      <b-navbar-item class="is-tab" tag="router-link" :to="{ name: 'Home' }">
        Week
      </b-navbar-item>

      <b-navbar-item class="is-tab"  tag="router-link" :to="{ name: 'Recipes' }">
        Recipes
      </b-navbar-item>

      <b-navbar-item tag="div" v-if="$route.name === 'Recipes'">
        <b-button class="is-primary" label="Add Recipe" @click="showAddRecipeForm = true"/>
        <b-modal v-model="showAddRecipeForm" scroll="keep">
          <AddRecipeForm @close="showAddRecipeForm = false"/>
        </b-modal>
      </b-navbar-item>
      <b-navbar-item tag="div" v-else>

        <b-button v-if="hasGaps" label="Fill gaps" class="is-primary" @click="onFillGaps"/>
        <b-button v-else label="Clear week" class="is-danger" @click="onClearWeek"/>

      </b-navbar-item>

    </template>
  </b-navbar>
</template>

<script>
import AddRecipeForm from '@/components/AddRecipeForm.vue';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'Navbar',
  components: {
    AddRecipeForm,
  },
  data: () => ({
    showAddRecipeForm: false,
  }),
  computed: {
    ...mapGetters([
      'hasGaps',
    ]),
  },
  methods: {
    ...mapActions(['fillGaps', 'emptyWeek']),
    onFillGaps() {
      this.fillGaps();
    },
    onClearWeek() {
      this.$buefy.dialog.confirm({
        message: 'Remove all recipes chosen for the week?',
        onConfirm: () => this.emptyWeek(),
      });
    },
  },
};
</script>

<style scoped>

</style>
