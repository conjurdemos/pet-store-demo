package hello.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
class VulnerableController {
  @GetMapping("/vulnerable")
  ResponseEntity<?> getEnvironment() {
      return ResponseEntity.ok().body(System.getenv());
  }
}
