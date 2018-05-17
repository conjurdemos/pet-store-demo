package hello.controller;

import hello.model.Pet;
import hello.repository.PetRepository;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/pets")
class PetCollectionController {
  @Autowired
  private PetRepository repository;

  @GetMapping()
  ResponseEntity<List<Pet>> getPets() {
    final List<Pet> petList = new ArrayList<>();
    final Iterable<Pet> all = repository.findAll();
    all.forEach(petList::add);
    return ResponseEntity.ok().body(petList);
  }
}