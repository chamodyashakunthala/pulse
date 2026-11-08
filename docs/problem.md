# <name> Incident Copilot

## 1. Who is it for?
A developer or small team running 3 to 5 backend services who can't
afford enterprise monitoring tools like Datadog.

## 2. What goes wrong for them today?
When something breaks, they search log files service by service.
That is slow, and they can't easily tell whether the problem began
right after a deploy or what changed.

## 3. What version 1 does
1. Collects events (service, latency, status, version) through an
   ingest API and a message queue, and stores them.
2. Detects abnormal behavior (a latency or error-rate spike compared
   to that service's recent normal) and raises an alert.
3. An AI agent reads the evidence (metrics and recent deploys) and
   writes a plain-language probable cause.

## 4. What version 1 does NOT do
- No full log search engine or distributed tracing
- No automatic fixes or rollbacks (it only explains)
- No user accounts, teams, or billing
- Not built for thousands of services
- Dashboard stays simple (charts and an incident list)

## 5. How I will know it works (starting targets)
- Handles 1,000 events/sec on my laptop with zero events lost
- Detects a simulated latency spike (80 ms to 900 ms) within 30 seconds
- The agent names the right cause in at least 8 of 10 simulated incidents
- Fewer than 1 false alarm per hour during normal simulated traffic
- Tests and CI pass, and the whole system starts with one
  `docker compose up`
