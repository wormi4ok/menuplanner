<template>
  <b-navbar class="mb-2">
    <template #brand>
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/flaningo.jpeg" alt="MenuPlanner logo"/>
      </b-navbar-item>
    </template>
    <template #start>
      <b-navbar-item>
        <a v-if="hasGaps" class="button is-primary" @click="onFillGaps">
          <strong>Fill gaps</strong>
        </a>
        <a v-else class="button is-danger" @click="onClearWeek">
          <strong>Clear week</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item>
        <b-button label="Recipes" tag="router-link" :to="{ path: '/recipes' }"/>
      </b-navbar-item>
      <b-navbar-item>
        <b-button label="Add Recipe" @click="showAddRecipeForm = true"/>
        <b-modal v-model="showAddRecipeForm" scroll="keep">
          <AddRecipeForm @close="showAddRecipeForm = false"/>
        </b-modal>
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
