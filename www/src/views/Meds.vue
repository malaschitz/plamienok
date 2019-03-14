<template>
  <v-container fluid>
    <v-toolbar
      flat
      color="white"
    >
      <v-toolbar-title>Lieky</v-toolbar-title>
      <v-spacer />
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
      <tr
        slot="items"
        slot-scope="props"
        @click="viewItem(props.item)"
      >
        <td class="text-xs-left">
          {{ props.item.Nazov }}
        </td>
        <td class="text-xs-left">
          {{ props.item.DoplnokNazvu }}
        </td>
      </tr>
    </v-data-table>

    <v-dialog
      v-model="dialog"
      max-width="500px"
      @keydown.esc="close"
      persistent
    >
      <v-card>
        <v-card-title>
          <span class="headline">Liek</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout
              column
              wrap
            >
              <v-text-field
                readonly
                label="Kód ŠÚKL"
                v-model="chosen.KodSukl"
              />
              <v-text-field
                readonly
                label="Registračné číslo"
                v-model="chosen.RegistracneCislo"
              />
              <v-text-field
                readonly
                label="Názov"
                v-model="chosen.Nazov"
              />
              <v-text-field
                readonly
                label="Doplnok Názvu"
                v-model="chosen.DoplnokNazvu"
              />
              <v-text-field
                readonly
                label="ATC"
                v-model="chosen.AtcText"
              />
              <v-text-field
                readonly
                label="Veľkosť balenia"
                v-model="chosen.VelkostBalenia"
              />
            </v-layout>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="blue darken-1"
            flat
            @click="close"
          >
            Zrušiť
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  name: "Meds",

  data: () => ({
    headers: [
      { text: "Názov", value: "Nazov", align: "left" },
      { text: "Doplnok názvu", value: "DoplnokNazvu" }
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
        .get("/api/meds?filter=" + this.filter)
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
    console.log("MEDS.VUE, mounted");
    this.readData();
  }
};
</script>

<style scoped>
</style>
