// Main file for dnd5engine

package main

import (
	"fmt"
)

func main() {
	fmt.Println("DnD 5E Engine Main Menu")
	
	/*	Block of code to generate ability modifiers
		Uses the equation abilityModifier = (baseStat-10)/2

		Sample Values:
		Ability		Modifier
		1			-5
		2-3			-4
		4-5 		-3
		6-7			-2
		8-9			-1
		10-11		+0
		12-13		+1
		14-15		+2
		etc...

		transfer these functions to charcterSheet struct when ready
	*/
	var baseStat, abilityModifier, mathCheck int

	mathCheck = baseStat - 10

	// Code rounds down, so this corrects mistakes with negative numbers
	if mathCheck < 0 {
		mathCheck--
	}

	abilityModifier = mathCheck / 2
	fmt.Println("baseStat: ", baseStat, "mathCheck: ", mathCheck, "abilityModifier: ", abilityModifier)
	
	baseStat = 19
	fmt.Println(baseStat % 18)
}

type characterSheet struct {
	
	// Base Ability Scores
	var strength int 
	var dexterity int 
	var constitution int
	var intelligence int  
	var wisdom int 
	var charisma int 
	
	// Hit Point Variables
	var hitPointMax int
	var hitPointCurrent int 
	var hitDiceType int
	var hitDiceMax int 
	var hitDiceCurrent int
	//var temporaryHitPoints int (disabled until I decide how to tackle this)
	
	// Defensive Scores
	var armorClass int
	var strengthSavingThrow int
	var dexteritySavingThrow int
	var constitutionSavingThrow int
	var intelligenceSavingThrow int
	var wisdomSavingThrow int
	var charismaSavingThrow int
	var deathSavingThrowSucess int
	var deathSavingThrowFailure int
	

}
