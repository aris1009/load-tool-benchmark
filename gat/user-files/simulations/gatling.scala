package poc

import io.gatling.core.Predef._
import io.gatling.http.Predef._

class BaseSimulation extends Simulation {
    val httpProtocol = http.baseUrl("http://127.0.0.1:3005")

    val scn = scenario("BasicSimulation")
        .exec(
            http("Homepage")
                .get("/")
                .check(status.is(200))
        )

    setUp(
        scn.inject(
            constantConcurrentUsers(15000).during(5 * 60)
        )
    ).protocols(httpProtocol)
}
