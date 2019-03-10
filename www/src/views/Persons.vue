<!--
  - Developed by Richard Malaschitz on 3/4/19 8:04 PM
  - Last modified 3/4/19 8:04 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-toolbar
      flat
      color="white"
    >
      <v-toolbar-title>Deti, príbuzní, iné osoby</v-toolbar-title>
      <v-spacer />
      <v-dialog
        v-model="dialog"
        max-width="500px"
      >
        <v-btn
          slot="activator"
          color="primary"
          dark
          class="mb-2"
        >
          Nový
        </v-btn>
        <v-card>
          <v-form
            ref="personform"
            v-model="valid"
            lazy-validation
          >
            <v-card-title>
              <span class="headline">Nové dieťa (klient)</span>
            </v-card-title>

            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex
                    xs12
                    sm8
                    md6
                  >
                    <v-text-field
                      v-model="editedItem.firstname"
                      label="Meno"
                      :rules="[rules.required]"
                      required
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm8
                    md6
                  >
                    <v-text-field
                      v-model="editedItem.surname"
                      label="Priezvisko"
                      :rules="[rules.required]"
                      required
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-menu
                      v-model="menuDatePicker"
                      :close-on-content-click="false"
                      :nudge-right="40"
                      lazy
                      transition="scale-transition"
                      offset-y
                      full-width
                      min-width="290px"
                    >
                      <template v-slot:activator="{ on }">
                        <v-text-field
                          v-model="editedItem.birthdate"
                          label="Dátum narodenia"
                          prepend-icon="event"
                          readonly
                          v-on="on"
                          :rules="[rules.required]"
                          required
                        />
                      </template>
                      <v-date-picker
                        v-model="editedItem.birthdate"
                        @input="menuDatePicker = false"
                      />
                    </v-menu>
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

    <v-layout>
      <v-flex xs12 md4 class="mr-1">
        <v-text-field v-model="filter.fulltext" label="Hľadať" @change="readData" />
      </v-flex>
      <v-flex xs12 md4 class="mr-1">
        <v-select :items="filterClients" v-model="filter.clients" label="Klienti" @change="readData" />
      </v-flex>
      <v-flex xs12 md4 class="mr-1">
        <v-select :items="filterOddelenie" v-model="filter.oddelenie" label="Oddelenie" @change="readData" />
      </v-flex>
      <v-flex xs12 md4 class="mr-1">
        <v-select :items="filterStav" v-model="filter.stav" label="Stav" @change="readData" />
      </v-flex>
    </v-layout>

    <v-data-table
      :headers="headers"
      :pagination.sync="pagination"
      :items="persons"
      class="elevation-1"
      rows-per-page-text="Počet riadkov"
    >
      <template
        slot="items"
        slot-scope="props"
      >
        <td class="text-xs-left">
          {{ props.item.FirstName }} {{ props.item.Surname }}
        </td>
        <td class="text-xs-right">
          {{ props.item.BirthDate.Day }}. {{ props.item.BirthDate.Month }}. {{ props.item.BirthDate.Year }}
        </td>
        <td class="text-xs-left">
          <v-chip
            v-if="props.item.IsHC"
            color="primary"
            text-color="white"
            small
          >
            HC
          </v-chip>
          <v-chip
            v-if="props.item.IsCGT"
            color="primary"
            text-color="white"
            small
          >
            CGT
          </v-chip>
          <v-chip
            v-if="props.item.Death"
            color="black"
            text-color="white"
            small
          >
            {{ props.item.Death.Month }}.{{ props.item.Death.Year }}
          </v-chip>
        </td>
        <td>
          {{ pocetDni( props.item) }}
        </td>
        <td class="justify-center layout px-0">
          <v-icon class="mr-2" @click="$router.push('/person/' + props.item.ID)">
            edit
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
    import '../helper.js';

    export default {
        name: "Persons",

        data: () => ({
          headers: [
            {text: 'Meno', value: 'Surname', align: 'left',  sortable: false},
            {text: 'Narodený', value: 'Birthdate', align: 'right', sortable: false},
            {text: 'Info', sortable: false},
            {text: 'Počet dní', sortable: false},
            {text: 'Akcie', sortable: false}],
          persons: [],
          pagination: {rowsPerPage:10},
          editedIndex: -1,
          editedItem: {
            firstname: '',
            surname: '',
            birthdate: '',
          },
          defaultItem: {
              firstname: '',
              surname: '',
              birthdate: '',
          },
          dialog: false,
          rules: {
            required: value => !!value || 'Required.',
          },
          valid: true,
          filter: {
            fulltext: '',
            clients: '',
            oddelenie: '',
            stav: '',
          },
          menuDatePicker: false,
          filterClients: [
                {text: 'Deti, klienti', value: 'd'},
                {text: 'Príbuzní', value: 'p'},
                {text: '-', value: 'a'},
            ],
          filterOddelenie: [
              {text: 'HomeCare', value: 'hc'},
              {text: 'CGT', value: 'cgt'},
              {text: '-', value: 'a'},
          ],
          filterStav: [
              {text: 'Prijatý v Plamienku', value: 'p'},
              {text: 'Odídený', value: 'o'},
              {text: 'Zomretý', value: 'z'},
              {text: '-', value: 'a'},
          ],
        }),

        computed: {

        },

        watch: {
          dialog (val) {
            val || this.close()
          }
        },

        methods: {
          close: function() {
            this.dialog = false
            setTimeout(() => {
              this.editedItem = Object.assign({}, this.defaultItem)
              this.editedIndex = -1
            }, 300)
          },

          save: function() {
            if (!this.$refs.personform.validate()) {
              return;
            }

            this.$axios.post("/r/api/person",this.editedItem).then(response => {
              if (response.status == 202) {
                this.$store.commit("alert","Chyba: " + response.data.Error);
              } else { //OK
                this.close()
                this.$router.push("/person/" + response.data.ID);
              }
            }).catch(response => {
              console.log('WRONG',response);
              this.$store.commit("alert","Chyba: " + response);
            });
          },

          readData: function() {
            this.$axios.post('/r/api/persons',this.filter).then(response => {
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
          },

          pocetDni: function(person) {
              if (person.PlamPrijatie) {
                  var firstDate = new Date(person.PlamPrijatie.Year,person.PlamPrijatie.Month,person.PlamPrijatie.Day);
                  var secondDate = new Date();
                  if (person.PlamPrepustenie) {
                      secondDate = new Date(person.PlamPrepustenie.Year,person.PlamPrepustenie.Month,person.PlamPrepustenie.Day);
                  }
                  var days = this.$root.days_between(firstDate, secondDate);
                  return days + " dní";
              } else {
                  return '-';
              }
          },

        },

        mounted: function() {
          console.log("mounted persons");
          this.readData();
        },


    }
</script>

<style scoped>

</style>