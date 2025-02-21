package main

import "archiver/internal/archiver"

var text = "Lorem ipsum dolor sit amet consectetur adipisicing elit. Sint ex adipisci exercitationem eaque voluptatem alias, fugiat ad quaerat earum recusandae, fuga tempora maiores aut corrupti. Et, eaque quod provident facere doloribus dicta reprehenderit aperiam mollitia quisquam fugiat molestiae consequuntur reiciendis repellendus voluptas, impedit explicabo ratione! Impedit iure illo vitae est distinctio ratione, commodi itaque, voluptatem iste eligendi esse ex, sed blanditiis quia voluptatibus autem fugiat nihil voluptate optio error eum quod harum corrupti. Quibusdam corporis consequatur provident tenetur, sequi cupiditate sed aliquid, quod blanditiis quidem molestiae! Nemo eius suscipit quos? Nostrum voluptatum accusantium ducimus magni ab autem. Tempore harum quidem animi consectetur! Est nobis sequi iure facere, ipsum natus cum nostrum rem voluptatum tempora, veniam, sapiente sint reprehenderit. Deleniti, ratione doloribus, amet ducimus harum molestias omnis sunt ut ea tempore illo hic aut facere, aperiam suscipit vel placeat quas exercitationem temporibus! Non accusantium impedit hic perferendis odit quis molestias aspernatur veritatis. Minima, ullam, sunt temporibus culpa consequuntur amet ea blanditiis consectetur vitae eaque officiis quis voluptates eveniet aperiam impedit voluptas debitis ipsa nostrum vel deleniti deserunt rerum. Temporibus similique expedita, vel, eaque ut quae illo animi officiis ea corrupti laudantium quia eos consequuntur necessitatibus delectus. Natus, atque. Obcaecati, ea laudantium!" // пока для тестирования будет текст, потом файлы

func main() {
	archiver.Huffman([]byte(text))
}
