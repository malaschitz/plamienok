<!--
  - Developed by Richard Malaschitz on 3/5/19 2:34 PM
  - Last modified 3/5/19 2:34 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <confirm ref="confirm" />
    <h2 class="headline">
      {{ person.FirstName }} {{ person.Surname }}
      <span class="subheading">{{ age }}</span>
    </h2>
    <v-tabs v-model="active" color="primary" dark>
      <v-tab key="0" ripple>
        Základné údaje
      </v-tab>
      <v-tab key="1" ripple>
        Zdravotné údaje
      </v-tab>
      <v-tab key="2" ripple>
        Návštevy, Telefonáty
      </v-tab>
      <v-tab key="4" ripple>
        Poradňa
      </v-tab>

      <v-tab-item key="0">
        <v-form v-model="valid" ref="baseform" lazy-validation>
          <v-container>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.FirstName"
                  :rules="nameRules"
                  :counter="10"
                  label="Meno"
                  required
                />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.Surname"
                  :rules="nameRules"
                  :counter="10"
                  label="Priezvisko"
                  required
                />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.RC"
                  label="Rodné číslo"
                  mask="######/####"
                />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsHC" label="HomeCare" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsCGT" label="CGT" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsPatient" label="Klient" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-select
                  v-model="person.Sex"
                  :items="[
                    { value: 0, text: '-' },
                    { value: 1, text: 'muž' },
                    { value: 2, text: 'žena' }
                  ]"
                  label=""
                />
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker1"
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
                      v-model="person.DtoBirthDate"
                      label="Dátum narodenia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker
                    v-model="person.DtoBirthDate"
                    @input="menuDatePicker1 = false"
                  />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker2"
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
                      v-model="person.DtoDeath"
                      label="Dátum úmrtia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker
                    v-model="person.DtoDeath"
                    @input="menuDatePicker2 = false"
                  />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker3"
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
                      v-model="person.DtoPlamPrijatie"
                      label="Dátum prijatia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker
                    v-model="person.DtoPlamPrijatie"
                    @input="menuDatePicker3 = false"
                  />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker4"
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
                      v-model="person.DtoPlamPrepustenie"
                      label="Dátum prepustenia"
                      prepend-icon="event"
                      v-on="on"
                      readonly
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker
                    v-model="person.DtoPlamPrepustenie"
                    @input="menuDatePicker4 = false"
                  />
                </v-menu>
              </v-flex>
            </v-layout>

            <h3>Kontaktné údaje</h3>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.Email"
                  :rules="emailRules"
                  label="E-mail"
                />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.Phone" label="Telefónne číslo" />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.Job" label="Zamestnanie" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.AddrStreet" label="Ulica" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.AddrCity" label="Mesto" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.AddrPSC"
                  label="PSČ"
                  mask="### ##"
                />
              </v-flex>
            </v-layout>
            <h3>Príbuzní</h3>

            <v-icon @click="relativeAdd">
              add
            </v-icon>
            <v-chip
              v-for="dr in person.DtoRelatives"
              label
              outline
              close
              @input="relativeClose(dr)"
              @click="relativeDetail(dr)"
              :key="dr.ID"
            >
              <v-icon left>
                person
              </v-icon>
              <strong>{{ dr.Relationship.Relation }}</strong>: {{ dr.Person.FirstName }} {{ dr.Person.Surname }}
            </v-chip>

            <v-layout>
              <v-btn color="info" @click="saveBase">
                Uložiť
              </v-btn>

              <v-btn color="warning" @click="readData">
                Refresh
              </v-btn>
            </v-layout>
          </v-container>
        </v-form>
      </v-tab-item>

      <v-tab-item key="1">
        <v-form v-model="valid" ref="zuform" lazy-validation>
          <v-container>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-select
                  v-model="person.ZP"
                  :items="[
                    { value: 'VZP', text: 'Všeobecná ZP' },
                    { value: 'dovera', text: 'Dôvera' },
                    { value: 'union', text: 'Union' }
                  ]"
                  label="Zdravotná poisťovňa"
                  clearable
                />
              </v-flex>
              <v-flex xs12 md8 sm6>
                <v-text-field
                  v-model="person.Odoslal"
                  label="Odosielajúci lekár"
                />
              </v-flex>
              <v-flex xs12 md12 sm12>
                <v-autocomplete
                  v-model="person.DGN"
                  :items="dgns"
                  chips
                  clearable
                  label="Zadanie diagnózy"
                  cache-items
                  hide-no-data
                  hide-details
                  multiple
                />
              </v-flex>
              <v-flex xs12 md12 sm12>
                <v-textarea
                  auto-grow
                  box
                  color="deep-purple"
                  rows="1"
                  v-model="person.DGNpopis"
                  label="Bližší popis diagnózy"
                />
              </v-flex>
              <v-flex xs12 md12 sm12>
                <v-textarea
                  auto-grow
                  box
                  color="deep-purple"
                  rows="1"
                  v-model="person.LekarPKontaktu"
                  label="Lekár prvého kontaku"
                />
              </v-flex>
              <v-flex xs12 md12 sm12>
                <v-textarea
                  auto-grow
                  box
                  color="deep-purple"
                  rows="1"
                  v-model="person.Alergia"
                  label="Popis alergií"
                />
              </v-flex>
              <v-flex xs12 md12 sm12>
                <v-textarea
                  auto-grow
                  box
                  color="deep-purple"
                  rows="1"
                  v-model="person.Laboratoria"
                  label="Kontakt na laboratóriá"
                />
              </v-flex>

              <v-btn color="info" @click="saveZP">
                Uložiť
              </v-btn>

              <v-btn color="warning" @click="readData">
                Refresh
              </v-btn>
            </v-layout>
          </v-container>
        </v-form>
      </v-tab-item>

      <v-tab-item key="2">
        <v-container fluid>
          <v-toolbar flat color="white">
            <v-btn @click="newVisit('H')">
              <v-icon>home</v-icon> Nová návšteva
            </v-btn>
            <v-btn @click="newVisit('P')">
              <v-icon>phone</v-icon> Nový telefonát
            </v-btn>
          </v-toolbar>
          <v-data-table
            :headers="[
              { text: 'Dátum', value: 'DtoDatum' },
              { text: 'Typ', value: 'DtoTyp' },
              { text: 'Popis', value: 'Popis' },
              { text: 'Detailný Popis', value: 'PopisDetail' },
              { text: 'Akcie', scrollable: false, value: 'ID' }
            ]"
            :items="visits"
            :items-per-page="20"
            class="elevation-1"
          >
            <template v-slot:items="props">
              <td class="text-xs-left">
                {{ props.item.DtoDatum }}
              </td>
              <td class="text-xs-left">
                <v-icon v-if="props.item.DtoTyp == 'H'">
                  home
                </v-icon>
                <v-icon v-if="props.item.DtoTyp == 'P'">
                  phone
                </v-icon>
              </td>
              <td class="text-xs-left text-no-wrap">
                {{ props.item.Popis }}
              </td>
              <td
                class="text-xs-left text-no-wrap text-truncate"
                style="max-width: 200px;"
              >
                {{ props.item.PopisDetail }}
              </td>
              <td>
                <v-icon class="mr-2" @click="editVisit(props.item)">
                  edit
                </v-icon>
              </td>
            </template>
          </v-data-table>
        </v-container>
      </v-tab-item>

      <v-tab-item key="4">
        <v-card flat>
          <v-card-text>Poradňa</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>

    <!-- Dialog - pridanie pribuzneho BEGIN -->

    <v-dialog v-model="dialogRelative" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Pridanie príbuzenského vzťahu</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12 sm12 md12 v-if="editRelative.msg">
                <span class="title red--text"
                  >Chyba: {{ editRelative.msg }}</span
                >
              </v-flex>
              <v-flex xs12 sm6 md4>
                <v-select
                  :items="relationships"
                  label="Príbuzenský vzťah"
                  v-model="editRelative.relationship"
                />
              </v-flex>
              <v-flex xs12 sm6 md8>
                <v-autocomplete
                  v-model="editRelative.person"
                  :items="persons"
                  label="Osoba"
                  clearable
                />
              </v-flex>
              <template v-if="!editRelative.person">
                <h3>Zadajte základné údaje ak sa jedná o novú osobu</h3>

                <v-flex xs12 sm6 md4>
                  <v-text-field label="Meno" v-model="editRelative.firstname" />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field
                    label="Priezvisko"
                    v-model="editRelative.surname"
                  />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field label="Telefón" v-model="editRelative.phone" />
                </v-flex>
              </template>
            </v-layout>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" flat @click="dialogRelative = false">
            Close
          </v-btn>
          <v-btn color="blue darken-1" flat @click="saveRelative()">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- Dialog - pridanie pribuzneho END -->
  </v-container>
</template>

<script>
import confirm from "../components/Confirm";

export default {
  name: "Person",

  components: { confirm: confirm },

  data: () => ({
    person: {
      DtoRelatives: []
    },
    visits: [],
    sessions: [],
    active: 0,
    valid: false,

    nameRules: [v => !!v || "Meno je vyžadované"],

    emailRules: [
      // eslint-disable-next-line
              v => ( /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || !v)  || 'E-mail must be valid'
    ],

    editRelative: {
      person: null,
      relationship: "",
      firstname: "",
      surname: "",
      phone: "",
      msg: ""
    },

    menuDatePicker1: false,
    menuDatePicker2: false,
    menuDatePicker3: false,
    menuDatePicker4: false,

    dgns: [],
    dialogRelative: false,
    relationships: [],
    persons: []
  }),

  methods: {
    readData: function() {
      this.$axios
        .get("/r/api/person/" + this.$route.params.id)
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.person = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });

      this.$axios
        .get("/r/api/visits/" + this.$route.params.id)
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.visits = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });

      this.$axios
        .get("/r/api/persons")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.persons = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },

    readCodebooks: function() {
      this.$axios
        .get("/api/diagnosesAll")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.dgns = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });

      this.$axios
        .get("/api/relationships")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.relationships = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },

    saveBase: function() {
      if (!this.$refs.baseform.validate()) {
        return;
      }
      this.$axios
        .put("/r/api/person", this.person)
        .then(response => {
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            //OK
            this.person = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },

    saveZP: function() {
      if (!this.$refs.zuform.validate()) {
        return;
      }
      this.$axios
        .put("/r/api/person", this.person)
        .then(response => {
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            //OK
            this.person = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },

    editVisit: function(item) {
      if (item.Typ == "H") {
        this.$router.push("/visitHome/" + item.ID);
      } else {
        this.$router.push("/visitPhone/" + item.ID);
      }
    },

    newVisit: function(typ) {
      if (typ == "H") {
        this.$router.push("/visitHome/new");
      } else {
        this.$router.push("/visitPhone/new");
      }
    },

    relativeDetail: function(dr) {
      this.$router.push("/person/" + dr.Person.ID);
    },

    relativeClose: function(dr) {
      this.$refs.confirm
        .open("Vymazať  vzťah ?", "Ste si istý ?", { color: "red" })
        .then(confirm => {
          console.log(confirm);
          this.$axios
            .delete("/r/api/person/relative/" + dr.ID)
            .then(response => {
              if (response.status == 202) {
                this.$store.commit("alert", "Chyba: " + response.data.Error);
              } else {
                this.readData();
              }
            })
            .catch(response => {
              this.$store.commit("alert", "Chyba: " + response);
            });
        });
    },

    relativeAdd: function() {
      this.editRelative.firstname = "";
      this.editRelative.surname = "";
      this.editRelative.msg = "";
      this.editRelative.phone = "";
      this.editRelative.relationship = "";
      this.dialogRelative = true;
    },

    saveRelative: function() {
      if (this.editRelative.relationship === "") {
        this.editRelative.msg = "Chýba zadaný príbuzenský vzťah";
        this.$store.commit("alert", this.editRelative.msg);
        return;
      }
      if (this.editRelative.person == null) {
        if (
          this.editRelative.firstname === "" ||
          this.editRelative.surname === ""
        )
          this.editRelative.msg =
            "Je nutné vybrať osobu zo zoznamu osôb alebo zadať novú";
        this.$store.commit("alert", this.editRelative.msg);
        return;
      }
      this.$axios
        .post("/r/api/person/relative/" + this.person.ID, this.editRelative)
        .then(response => {
          if (response.status == 202) {
            this.editRelative.msg = response.data.Error;
            this.$store.commit("alert", this.editRelative.msg);
          } else {
            this.dialogRelative = false;
            this.readData();
          }
        })
        .catch(response => {
          this.editRelative.msg = response;
          this.$store.commit("alert", this.editRelative.msg);
        });
    }
  },

  mounted: function() {
    console.log("mounted person", this.$route.params.id);
    this.readCodebooks();
    this.readData();
  },

  computed: {
    age: function() {
      var dob = this.person.DtoBirthDate;
      if (dob && dob.length === 10) {
        var year = Number(dob.substr(0, 4));
        var month = Number(dob.substr(4, 2)) - 1;
        var day = Number(dob.substr(6, 2));
        var today = new Date();
        var age = today.getFullYear() - year;
        if (
          today.getMonth() < month ||
          (today.getMonth() == month && today.getDate() < day)
        ) {
          age--;
        }
        return "(" + age + "r)";
      } else {
        return "";
      }
    }
  },

  watch: {
    $route() {
      this.readData();
    }
  }
};
</script>

<style scoped></style>
