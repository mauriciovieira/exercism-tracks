class Clock
  attr_reader :raw_hour, :raw_minute

  def initialize(params)
    @raw_hour = params[:hour] || 0
    @raw_minute = params[:minute] || 0
  end

  def to_s
    "%02d:%02d" % [hour, minute]
  end

  def +(other)
    Clock.new(hour: @raw_hour + other.raw_hour, minute: @raw_minute + other.raw_minute)
  end

  def -(other)
    Clock.new(hour: @raw_hour - other.raw_hour, minute: @raw_minute - other.raw_minute)
  end

  def ==(other)
    hour == other.hour && minute == other.minute
  end

  def hour
    (@raw_hour + @raw_minute / 60) % 24
  end

  def minute
    @raw_minute % 60
  end
end