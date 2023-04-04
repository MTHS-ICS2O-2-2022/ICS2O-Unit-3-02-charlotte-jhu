// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

/**
 * This function calculates the volume of a square-based pyramid
 */

function myButtonClicked() {
  // input
  const length = parseFloat(document.getElementById('Length').value)
  const width = parseFloat(document.getElementById('Width').value)
  const height = parseFloat(document.getElementById('Height').value)

  // process
  const volume = (length * width * height) / 3

  // output
  document.getElementById('volume').innerHTML = 'Volume is ' + volume + 'mm³'
}
