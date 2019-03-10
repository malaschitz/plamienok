<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Ciele</v-toolbar-title>
    </v-toolbar>
    <v-list dense>
      <template v-for="(person) in persons">
        <h3 :key="'nadpis' + person.value">
          {{ person.text }}
        </h3>
        <PersonTasks :key="'tasks' + person.value" :person-id="person.value" />
      </template>
    </v-list>
  </v-container>
</template>

<script>
    import '../helper.js';
    import PersonTasks from './PersonTasks.vue'


    export default {
        name: "Tasks",
        components: {
          PersonTasks
        },

        data: () => ({
          persons: [],
        }),

        computed: {

        },

        watch: {
        },

        methods: {
          readData: function() {
            this.$axios.get('/r/api/persons/tasks').then(response => {
              console.log('OK',response);
              if (response.status == 202) {
                this.$store.commit("alert","Chyba: " + response.data.Error);
              } else {
                this.persons = response.data;
              }
            }).catch(response => {
              console.log('WRONG',response);
              this.$store.commit("alert","Chyba: " + response);
            });
          }
        },

        mounted: function() {
          console.log("mounted persons");
          this.readData();
        },
    }
</script>

<style scoped>

</style>