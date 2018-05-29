package hello.controller;

import java.net.URI;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.http.ResponseEntity;

@RestController
class VulnerableController {
  @GetMapping("/vulnerable")
  ResponseEntity<?> getEnvironment() {
      return ResponseEntity.ok().body(System.getenv());
  }
}
