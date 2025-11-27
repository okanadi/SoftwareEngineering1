<script lang="ts">
let name = "Max";
let age = 23;
let warn: string | null = null;

const increaseAge = () => age += 1;
const decreaseAge = () => age -= 1;

$:uppercase = name.toUpperCase();
$:if (age > 40) {
    warn = "zu alt";
} else {
    warn = "alles gut"
}
</script>

<main>
    <p>{name}</p>
    <p>{age}</p>
    <p>{uppercase}</p>
    <p>{warn}</p>
    <button on:click={increaseAge}>Increase</button>
    <button on:click={decreaseAge}>decrease</button>
</main>

