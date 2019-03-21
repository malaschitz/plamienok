<!--
  - Developed by Richard Malaschitz on 3/4/19 1:25 PM
  - Last modified 3/4/19 1:16 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Služobné autá</v-toolbar-title>
      <v-divider class="mx-2" inset vertical />
      <v-spacer />
      <v-checkbox v-model="deleted" label="Zrušené autá" />
      <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" dark class="mb-2">
          Nové auto
        </v-btn>
        <v-card>
          <v-form ref="carform" v-model="valid" lazy-validation>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex xs12 sm8 md6>
                    <v-text-field
                      v-model="editedItem.Name"
                      label="Názov"
                      :rules="[rules.required, rules.max20]"
                      required
                    />
                  </v-flex>
                  <v-flex xs12 sm12 md12>
                    <v-textarea
                      v-model="editedItem.Popis"
                      label="Popis"
                      hint="Popis auta, jeho výbava a iné dôležité informácie"
                    />
                  </v-flex>
                </v-layout>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" flat @click="close">
                Zrušiť
              </v-btn>
              <v-btn color="blue darken-1" flat @click="save">
                Uložiť
              </v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-dialog>
    </v-toolbar>

    <v-data-table
      :headers="headers"
      :pagination.sync="pagination"
      :items="filteredCars"
      class="elevation-1"
      rows-per-page-text="Počet riadkov"
    >
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">
          {{ props.item.Name }}
        </td>
        <td class="text-xs-left">
          {{ props.item.Popis }}
        </td>
        <td class="justify-center layout px-0">
          <v-icon small class="mr-2" @click="editItem(props.item)">
            edit
          </v-icon>
          <v-icon v-if="!deleted" small @click="deleteItem(props.item)">
            delete
          </v-icon>
          <v-icon v-else small @click="deleteItem(props.item)">
            autorenew
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
export default {
  name: "Cars",

  data: () => ({
    headers: [
      { text: "Skratka", value: "Name", align: "left" },
      { text: "Popis", value: "Popis" },
      { text: "Akcie", sortable: false }
    ],
    cars: [],
    pagination: { rowsPerPage: 10 },
    editedIndex: -1,
    editedItem: {
      ID: "",
      Name: "",
      Popis: ""
    },
    dialog: false,
    rules: {
      required: value => !!value || "Required.",
      max20: value => (value || "").length <= 20 || "Max 20 characters"
    },
    valid: true,
    deleted: false
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Nové auto" : "Editácia auta";
    },

    filteredCars() {
      var a = [];
      for (var i = 0; i < this.cars.length; i++) {
        var u = this.cars[i];
        if ((this.deleted && u.Deleted) || (!this.deleted && !u.Deleted)) {
          a.push(u);
        }
      }
      return a;
    }
  },

  watch: {
    dialog(val) {
      val || this.close();
    }
  },

  methods: {
    editItem: function(item) {
      this.editedIndex = this.cars.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    deleteItem: function(item) {
      var conf;
      if (!this.deleted) {
        conf = "Ste si istý, že chcete vymazať auto ?";
      } else {
        conf = "Ste si istý, že chcete obnoviť auto ?";
      }
      if (confirm(conf)) {
        this.$axios
          .delete("/r/api/car/" + item.ID)
          .then(response => {
            if (response.status == 202) {
              this.$store.commit("alert", "Chyba: " + response.data.Error);
            } else {
              //OK
              this.readData();
            }
          })
          .catch(response => {
            console.log("WRONG", response);
            this.$store.commit("alert", "Chyba: " + response);
          });
      }
    },

    close: function() {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      }, 300);
    },

    save: function() {
      if (!this.$refs.carform.validate()) {
        return;
      }

      this.$axios
        .post("/r/api/car", this.editedItem)
        .then(response => {
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            //OK
            this.editedItem = response.data;

            if (this.editedIndex > -1) {
              Object.assign(this.cars[this.editedIndex], this.editedItem);
            } else {
              this.cars.push(this.editedItem);
            }

            this.close();
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },

    readData: function() {
      this.$axios
        .get("/r/api/cars")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.cars = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    }
  },

  mounted: function() {
    console.log("mounted cars");
    this.readData();
  }
};
</script>

<style scoped></style>
