package hello;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import javax.sql.DataSource;

import org.springframework.boot.autoconfigure.jdbc.DataSourceBuilder;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Component;

@Configuration
@Component
public class DataSourceConfig {

    @ConfigurationProperties(prefix = "spring.datasource")
    @Bean
    @Primary
    public DataSource getDataSource() throws IOException {
        String username, password;
        if(useSecretsFromFile()) {
            username = getSecretFromFile(Paths.get(secretsDirectory(), "DB_USERNAME"));
            password = getSecretFromFile(Paths.get(secretsDirectory(), "DB_PASSWORD"));
        } else
        {
            username = System.getenv("DB_USERNAME");
            password = System.getenv("DB_PASSWORD");
        }
        
        return DataSourceBuilder
                .create()
                .username(username)
                .password(password)
                .build();
    }

    private String getSecretFromFile(Path path) throws IOException {
        byte[] encoded = Files.readAllBytes(path);
        return new String(encoded, StandardCharsets.UTF_8).trim();
    }

    private boolean useSecretsFromFile() {
        String useFileSecrets = System.getenv("USE_FILE_SECRETS");
        return useFileSecrets != null && useFileSecrets.equalsIgnoreCase("true");
    }

    private String secretsDirectory() {
        String secretsDirectory = System.getenv("SECRETS_DIRECTORY");
        if(secretsDirectory == null || secretsDirectory.isEmpty()) {
            secretsDirectory = "/run/conjur/secrets";
        }

        return secretsDirectory;
    }
}