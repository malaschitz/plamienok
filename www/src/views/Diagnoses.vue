<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Diagnózy</v-toolbar-title>
      <v-spacer/>
      <v-text-field
        v-model="filter"
        flat
        hide-details
        label="Hladať"
        prepend-inner-icon="search"
        @change="readData"
      />
    </v-toolbar>

    <v-data-table
      :headers="headers"
      :pagination.sync="pagination"
      :items="items"
      class="elevation-1"
      rows-per-page-text="Počet riadkov"
    >
      <tr slot="items" slot-scope="props">
        <td class="text-xs-left">{{ props.item.Skratka }}</td>
        <td class="text-xs-left">{{ props.item.Popis }}</td>
      </tr>
    </v-data-table>

  </v-container>
</template>

<script>
export default {
  name: "Diagnoses",

  data: () => ({
    headers: [
      { text: "Skratka", value: "Skratka", align: "left" },
      { text: "Popis", value: "Popis" }
    ],
    items: [],
    pagination: { rowsPerPage: 10 },
    chosen: {},
    dialog: false,
    filter: ""
  }),

  computed: {
  },

  methods: {
    viewItem: function(item) {
      this.chosen = Object.assign({}, item);
      this.dialog = true;
    },
    close: function() {
      this.dialog = false
    },
    readData: function() {
      this.$axios
        .get("/api/diagnoses?filter=" + this.filter)
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.items = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    }
  },

  mounted: function() {
    console.log("DIAGNOSES.VUE, mounted");
    this.readData();
  }
};
</script>

<style scoped>
</style>
